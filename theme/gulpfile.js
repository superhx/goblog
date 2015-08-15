var gulp = require('gulp');
var rename = require('gulp-rename');
var uglify = require('gulp-uglify');
var shell = require('gulp-shell');
var minifyCSS = require('gulp-minify-css');
var react = require('gulp-react');


var DATA_JS_SRC = 'src/data/js/';
var DATA_JS_DEST = 'data/js/';
var DATA_CSS_SRC = 'src/data/css/';
var DATA_CSS_DEST = 'data/css';
var TMPL = 'template/'
// the blog dir for test
var BLOG_DIR = '/Users/Lancer/Desktop/blog/';

gulp.task('js', function () {
    gulp.src(DATA_JS_SRC + '*.js')
        .pipe(uglify())
        .pipe(rename({extname: '.min.js'}))
        .pipe(gulp.dest(DATA_JS_DEST))
})

gulp.task('jsx', function () {
    gulp.src(DATA_JS_SRC + '*.jsx')
        .pipe(react())
        .pipe(uglify())
        .pipe(rename({extname: '.min.js'}))
        .pipe(gulp.dest(DATA_JS_DEST))
})

gulp.task('css', function () {
    gulp.src(DATA_CSS_SRC + '*.css')
        .pipe(minifyCSS())
        .pipe(rename({extname: '.min.css'}))
        .pipe(gulp.dest(DATA_CSS_DEST))
})

gulp.task('template', shell.task([
    'rm -f category.json',
    'goblog g'
], {'cwd': BLOG_DIR}))

gulp.task('init', blogReloadResource)

gulp.task('monitor', function () {
    gulp.watch(DATA_JS_SRC + '*.js', ['js', 'init'])
    gulp.watch(DATA_JS_SRC + '*.jsx', ['jsx', 'init'])
    gulp.watch(DATA_CSS_SRC + '*.css', ['css', 'init'])
    gulp.watch(TMPL + '*.htm', ['template'])
})

gulp.task('default', ['js', 'jsx', 'css', 'template', 'init', 'monitor'])

function blogReloadResource() {
    gulp.src('').pipe(shell([
        'goblog i'
    ], {'cwd': BLOG_DIR}));
}